/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    serverActions: true,
  },
    async rewrites() {
        return [
          {
            source: '/api/:path*',
            destination: 'http://localhost:9001/:path*' // Proxy to Backend
          },
        ]
      }
}

module.exports = nextConfig
