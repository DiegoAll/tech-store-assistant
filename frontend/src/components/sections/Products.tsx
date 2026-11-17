import { motion } from 'motion/react'

const products = [
  { 
    id: 1, 
    name: 'SmartHome Hub X1', 
    price: 150, 
    oldPrice: 199, 
    discount: '-25%',
    rating: 4.8,
    image: 'https://images.unsplash.com/photo-1558002038-1055907df827?auto=format&fit=crop&q=80&w=400', 
    desc: 'Centro de control para tu hogar.' 
  },
  { 
    id: 2, 
    name: 'Cámara SecureCam Pro', 
    price: 89, 
    oldPrice: 120, 
    discount: '-15%',
    rating: 4.5,
    image: 'https://images.unsplash.com/photo-1557324232-b8917d3c3dcb?auto=format&fit=crop&q=80&w=400', 
    desc: 'Vigilancia 24/7 en HD con IA.' 
  },
  { 
    id: 3, 
    name: 'Termostato Smart', 
    price: 120, 
    oldPrice: null, 
    discount: null,
    rating: 4.9,
    image: 'https://images.unsplash.com/photo-1567924675639-6bc0dde5a4a5?auto=format&fit=crop&q=80&w=400', 
    desc: 'Ahorra energía automáticamente.' 
  },
]

export default function Products() {
  return (
    <section id="productos" className="py-20 bg-gray-950">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        
        <div className="flex justify-between items-end mb-10">
          <div>
            <h2 className="text-3xl font-bold text-white">Los más vendidos</h2>
            <p className="text-gray-400 mt-2">Tecnología diseñada para ti.</p>
          </div>
          <a href="#" className="text-blue-400 hover:text-blue-300 font-medium transition-colors">
            Ver todo el catálogo →
          </a>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
          {products.map((product, idx) => (
            <motion.div 
              key={product.id}
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ duration: 0.5, delay: idx * 0.1 }}
              className="bg-white rounded-xl overflow-hidden hover:shadow-2xl hover:shadow-blue-500/20 transition-all duration-300 group cursor-pointer relative"
            >
              {/* Badge de descuento */}
              {product.discount && (
                <span className="absolute top-3 left-3 bg-red-500 text-white text-xs font-bold px-2 py-1 rounded z-10">
                  {product.discount}
                </span>
              )}

              {/* Imagen */}
              <div className="h-56 overflow-hidden bg-gray-100 p-4 flex justify-center items-center">
                <img src={product.image} alt={product.name} className="h-full object-contain group-hover:scale-110 transition-transform duration-500" />
              </div>

              {/* Info */}
              <div className="p-5">
                <h3 className="text-lg font-bold text-gray-800 line-clamp-1">{product.name}</h3>
                
                {/* Rating */}
                <div className="flex items-center gap-1 mt-1">
                  <span className="text-yellow-500 text-sm">★</span>
                  <span className="text-gray-500 text-sm">{product.rating} (120 opiniones)</span>
                </div>

                {/* Precios */}
                <div className="mt-3">
                  {product.oldPrice && (
                    <span className="text-gray-400 text-sm line-through block">${product.oldPrice} USD</span>
                  )}
                  <span className="text-2xl font-extrabold text-gray-900">${product.price} USD</span>
                </div>

                {/* Botón */}
                <button className="mt-4 w-full bg-gray-900 hover:bg-blue-600 text-white font-semibold py-2 rounded-lg transition-colors flex items-center justify-center gap-2">
                  🛒 Añadir al carrito
                </button>
              </div>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  )
}