import type { Metadata } from 'next'
import './globals.css'
import Navbar from './components/Navbar'

export const metadata: Metadata = {
  title: 'EduHub - Nen tang Hoc tap & Ho tro Tam ly',
  description: 'Nen tang giao duc tich hop video hoc tap, am thanh song nao, chatbot tam ly Buddy AI va ma QR thong minh cho hoc sinh',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="vi">
      <body className="antialiased bg-gray-50 min-h-screen">
        <Navbar />
        <main className="min-h-screen">
          {children}
        </main>
        <footer className="bg-gray-900 text-gray-400 text-center py-6 mt-16">
          <p className="text-sm">© 2026 EduHub - Nen tang Hoc tap & Ho tro Tam ly Hoc sinh</p>
          <p className="text-xs mt-1">Duoc xay dung de ho tro su phat trien toan dien cua hoc sinh</p>
        </footer>
      </body>
    </html>
  )
}
