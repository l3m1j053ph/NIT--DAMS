document.getElementById('exploreBtn').addEventListener('click', function () {
    alert('Welcome to the Future!');
});

// Example of a simple lightbox functionality (optional enhancement)
const galleryItems = document.querySelectorAll('.gallery-item');

galleryItems.forEach(item => {
    item.addEventListener('click', () => {
        alert('This can be expanded into a full-screen lightbox viewer!');
    });
});
