import './globals.css'

export const metadata = {
  title: 'Twitcher',
  description: 'Automated music stream',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className='bg-background w-screen h-screen'>{children}</body>
    </html>
  )
}
