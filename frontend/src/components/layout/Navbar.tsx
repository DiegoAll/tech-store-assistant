import { motion } from 'motion/react'

export default function Navbar() {
  return (
    <motion.nav 
      initial={{ y: -100 }}
      animate={{ y: 0 }}
      transition={{ type: "spring", stiffness: 100 }}
      className="fixed top-0 w-full z-50 bg-gray-900 border-b border-gray-800"
    >
      {/* Barra superior */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-20">
          
          {/* Logo */}
          <div className="flex-shrink-0 flex items-center gap-2 cursor-pointer">
            <span className="text-3xl">🛍️</span>
            <span className="font-bold text-2xl text-white tracking-tight">TechStore</span>
          </div>
          
          {/* Barra de Búsqueda */}
          <div className="hidden md:flex flex-1 max-w-xl mx-8">
            <div className="relative w-full">
              <input 
                type="text" 
                placeholder="Buscar productos, marcas y más..." 
                className="w-full bg-white text-gray-900 rounded-full py-2 px-5 focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
              <button className="absolute right-0 top-0 h-full bg-blue-600 hover:bg-blue-700 text-white px-5 rounded-r-full transition-colors">
                🔍
              </button>
            </div>
          </div>

          {/* Iconos de Usuario y Carrito */}
          <div className="flex items-center gap-6">
            <button className="flex flex-col items-center text-gray-300 hover:text-white transition-colors">
              <span className="text-xl">👤</span>
              <span className="text-xs mt-1">Mi Cuenta</span>
            </button>
            <button className="flex flex-col items-center text-gray-300 hover:text-white transition-colors relative">
              <span className="text-xl">🛒</span>
              <span className="text-xs mt-1">Carrito</span>
              {/* Badge de cantidad */}
              <span className="absolute top-0 right-0 inline-flex items-center justify-center px-2 py-1 text-xs font-bold leading-none text-white transform translate-x-1/2 -translate-y-1/2 bg-red-600 rounded-full">
                3
              </span>
            </button>
          </div>
        </div>
      </div>
      
      {/* Barra de Navegación Secundaria (Categorías) */}
      <div className="bg-gray-800 border-t border-gray-700">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex items-center gap-6 h-12 text-sm font-medium text-gray-300">
            <a href="#" className="hover:text-white transition-colors flex items-center gap-2">☰ Todas las categorías</a>
            <a href="#" className="hover:text-white transition-colors">Tecnología</a>
            <a href="#" className="hover:text-white transition-colors">Hogar Inteligente</a>
            <a href="#" className="hover:text-white transition-colors">Ofertas</a>
            <a href="#" className="text-red-400 hover:text-red-300 transition-colors font-bold">Hot Sale 🔥</a>
          </div>
        </div>
      </div>
    </motion.nav>
  )
}