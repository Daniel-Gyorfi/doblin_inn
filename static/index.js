const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
        console.debug("entry exists")
        if (entry.isIntersecting) {
            console.debug("entry detected")
            entry.target.classList.add("reveal")
        } else {
            console.debug()
        }
    })
});

const fadeElems = document.querySelectorAll('.fade-in');
fadeElems.forEach((el) => observer.observe(el));