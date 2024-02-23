fetch("/schools/")
    .then((e) => e.json())
    .then((s) => {
    for (school in s) {
        document.querySelector("select").innerHTML += `<option value="${s[school]}">${s[school]}</option>`
    }
})
