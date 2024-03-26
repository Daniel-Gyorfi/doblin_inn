const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
        if (entry.isIntersecting) {
            entry.target.classList.add("reveal");
        }  
    })
});

const fadeElems = document.querySelectorAll('.fade-in');
fadeElems.forEach((el) => observer.observe(el));