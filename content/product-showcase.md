---
title: Premium Coffee Collection
date: 2024-02-01
template: product-grid
styles:
  primary-color: "#8B4513"
  accent-color: "#D4AF37"
  background: "#FFF8E7"
products:
  - name: Ethiopian Yirgacheffe
    price: 24.99
    stock: 150
    rating: 4.8
    description: "A light-bodied coffee with floral and citrus notes, perfect for those who enjoy complex flavors."
    image: "ethiopian-yirgacheffe.jpg"
  - name: Colombian Supremo
    price: 22.99
    stock: 200
    rating: 4.7
    description: "Rich and balanced with hints of caramel and chocolate, this coffee is a true classic."
    image: "colombian-supremo.jpg"
featured_image: coffee-beans-hero.jpg
---

<div style="background-color: #FFF8E7; font-family: Arial, sans-serif; padding: 2em;">
  <header style="text-align: center;">
    <img src="coffee-beans-hero.jpg" alt="Premium Coffee" style="max-width: 100%; height: auto; border-radius: 10px;">
    <h1 style="color: #8B4513; font-size: 3em;">Artisanal Coffee Selection</h1>
    <p style="color: #D4AF37; font-size: 1.2em;">Discover the finest coffee from around the world</p>
  </header>

  <section class="featured-section" style="margin: 2em 0; padding: 2em; background-color: #F3E5D5; border: 2px solid #D4AF37; border-radius: 10px;">
    <h2 style="color: #8B4513;">This Month's Specials</h2>

    <div class="product-grid" style="display: grid; grid-template-columns: repeat(auto-fill, minmax(250px, 1fr)); gap: 1.5em; margin-top: 1em;">
      {{range .products}}
      <div class="product-card" style="border: 1px solid #8B4513; border-radius: 10px; padding: 1em; background-color: #FFF8E7;">
        <img src="{{.image}}" alt="{{.name}}" style="width: 100%; border-radius: 10px;">
        <h3 style="color: #8B4513; margin: 0.5em 0;">{{.name}}</h3>
        <p style="color: #555; font-size: 0.9em;">{{.description}}</p>
        <p class="price" style="color: #D4AF37; font-weight: bold; font-size: 1.2em;">${{.price}}</p>
        <div class="stock-indicator" style="color: {{if gt .stock 100}}green{{else}}red{{end}}; font-size: 0.9em;">
          Stock: {{.stock}} units
        </div>
        <div class="rating" style="margin: 0.5em 0; font-size: 0.9em; color: #FFD700;">★ {{.rating}}/5.0</div>
        <button style="background-color: #8B4513; color: #FFF; border: none; padding: 0.5em 1em; border-radius: 5px; cursor: pointer;">Add to Cart</button>
      </div>
      {{end}}
    </div>
  </section>

  <section class="brewing-guide" style="margin-top: 2em; padding: 2em; background-color: #F8F1E8; border: 2px dashed #8B4513; border-radius: 10px;">
    <h2 style="color: #8B4513;">Brewing Guide</h2>
    <p style="color: #555; font-size: 1em;">Perfect your cup of coffee with these expert recommendations:</p>
    <ul style="color: #8B4513; line-height: 1.5em;">
      <li><strong>Temperature:</strong> 92°C</li>
      <li><strong>Grind Size:</strong> Medium-fine</li>
      <li><strong>Brew Time:</strong> 4 minutes</li>
    </ul>
    <button style="margin-top: 1em; background-color: #D4AF37; color: #8B4513; border: none; padding: 0.5em 1em; border-radius: 5px; cursor: pointer;">Download Full Guide</button>
  </section>

  <footer style="margin-top: 2em; text-align: center; padding: 1em; background-color: #8B4513; color: #FFF; border-radius: 5px;">
    <p>&copy; 2024 Premium Coffee Collection | Crafted with care for coffee lovers</p>
  </footer>
</div>
