document.addEventListener("DOMContentLoaded", function() {
    console.log("DAMS frontend loaded");

    // Smooth scroll for the Get Started button
    document.querySelector('.cta-button').addEventListener('click', function() {
        window.scrollTo({ top: document.body.scrollHeight, behavior: 'smooth' });
    });

    // Interactive feature card hover effects
    document.querySelectorAll('.feature-card').forEach(card => {
        card.addEventListener('mouseover', function() {
            card.style.transform = 'scale(1.05)';
        });
        card.addEventListener('mouseout', function() {
            card.style.transform = 'scale(1)';
        });
    });
});
