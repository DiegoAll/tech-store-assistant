import { motion } from 'motion/react'

export default function Hero() {
  return (
    <section className="relative bg-gray-900 pt-32 pb-20 lg:pt-48 lg:pb-32 overflow-hidden">
      {/* Efecto de fondo */}
      <div className="absolute inset-0 bg-[url('https://grainy-gradients.vercel.app/noise.svg')] opacity-20"></div>
      <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[600px] bg-blue-600/20 blur-[120px] rounded-full"></div>

      <div className="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <motion.h1 
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="text-4xl sm:text-6xl font-extrabold text-white tracking-tight mb-6"
        >
          El futuro de tu hogar, <br />
          <span className="text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-purple-500">
            en tus manos.
          </span>
        </motion.h1>
        
        <motion.p 
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 0.1 }}
          className="mt-4 text-xl text-gray-400 max-w-2xl mx-auto mb-10"
        >
          Descubre nuestra nueva línea de dispositivos SmartHome. Automatiza, protege y simplifica tu vida con tecnología de punta.
        </motion.p>

        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 0.2 }}
        >
          <button className="bg-blue-600 hover:bg-blue-700 text-white px-8 py-3 rounded-full font-semibold text-lg transition-all hover:scale-105 shadow-lg shadow-blue-600/30">
            Ver Catálogo
          </button>
        </motion.div>
      </div>
    </section>
  )
}