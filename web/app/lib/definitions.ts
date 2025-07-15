import { z } from 'zod'

export const RegisterFormSchema = z.object({
  name: z
    .string()
    .min(3, "Name must be at least 3 characters long.")
    .max(30, "Name should have 30 letters at max.")
    .trim(),
  email: z
    .string()
    .email("Please enter a valid email."),
  password: z
    .string()
    .min(8, { message: 'Be at least 8 characters long' })
    .regex(/[a-zA-Z]/, { message: 'Contain at least one letter.' })
    .regex(/[0-9]/, { message: 'Contain at least one number.' })
    .regex(/[^a-zA-Z0-9]/, {
      message: 'Contain at least one special character.',
    })
    .trim(),
})

export type RegisterFormSchemaType = z.infer<typeof RegisterFormSchema>;

export type FormState =
  | {
    data?: RegisterFormSchemaType,
    errors?: {
      name?: string[]
      email?: string[]
      password?: string[]
    }
    message?: string
  }
  | undefined
