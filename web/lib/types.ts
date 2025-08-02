import { z } from 'zod'

export const registerSchema = z.object({
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
    .min(8, { message: 'Password Be at least 8 characters long' })
    .regex(/[a-zA-Z]/, { message: 'Password must contain at least one letter.' })
    .regex(/[0-9]/, { message: 'Password must contain at least one number.' })
    .regex(/[^a-zA-Z0-9]/, {
      message: 'Password must contain at least one special character.',
    })
    .trim(),
})

export const loginSchema = z.object({
  email: z
    .string()
    .email("Please enter a valid email."),
  password: z
    .string()
    .trim(),
})

export const websiteSchema = z.object({
  url: z
    .string()
    .url("Please enter a valid URL.")
    .trim(),
  frequency: z
    .string()
    .min(1, "Frequency is required."),
  regions: z
    .array(z.string())
    .min(1, "At least one region must be selected."),
})
