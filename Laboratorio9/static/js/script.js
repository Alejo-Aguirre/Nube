document.addEventListener('DOMContentLoaded', function() {
    console.log('Página cargada correctamente');
    
    // Puedes añadir más funcionalidad JavaScript aquí
    document.querySelectorAll('.image-item').forEach(item => {
        item.addEventListener('click', function() {
            console.log('Imagen seleccionada:', this.textContent);
        });
    });
});