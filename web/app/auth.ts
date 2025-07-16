import axios, { AxiosError } from "axios";
import NextAuth from "next-auth";
import Credentials from "next-auth/providers/credentials";
import { API_URL } from "./constants";

type user = {
  id: string;
  name: string;
  email: string;
  avatar: string;
  createdAt: Date;
  updatedAt: Date;
}

declare module "next-auth" {
  interface Session {
    token: string;
    user: user;
  }
  interface User {
    token: string;
    user: user
  }
}

export const { handlers, signIn, signOut, auth } = NextAuth({
  session: {
    strategy: 'jwt',
  },
  providers: [
    Credentials({
      credentials: {
        email: {},
        password: {}
      },
      async authorize(credentials) {
        try {
          const response = await axios.post(`${API_URL}/auth/login`, {
            email: credentials.email,
            password: credentials.password
          });

          if (response.data && response.data.token) {
            return response.data
          }

          throw new Error('Login failed. No token received.');
        } catch (error) {
          if (error instanceof AxiosError) {
            throw new Error(error.response?.data?.error || "An error occurred during login.");
          }

          throw error;
        }
      },
    })
  ],
  callbacks: {
    async jwt({ token, user }) {
      if (user) {
        token.token = user.token
        token.user = user.user
      }

      return token;
    },
    async session({ session, token }) {
      session.token = token.token as string
      //@ts-ignore
      session.user = token.user as user

      return session;
    },
  },
  pages: {
    signIn: '/login',
  },
});
