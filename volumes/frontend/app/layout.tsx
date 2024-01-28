import type { Metadata } from 'next'

import './globals.css'
import ToastProvider from '@/lib/toast-provider'

export const metadata: Metadata = {
  title: 'Create Next App',
  description: 'Generated by create next app',
}

// SSG disabled
export const dynamic = 'force-dynamic'

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang='ja'>
      <body>
        <ToastProvider>{children}</ToastProvider>
      </body>
    </html>
  )
}
