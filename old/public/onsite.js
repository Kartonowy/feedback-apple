const checkbox = document.querySelector("#agreement");

checkbox.addEventListener("change", () => {
  document.querySelector("#send").disabled = !checkbox.checked;
});


document.querySelector("#send").addEventListener("click", () => {
    let kind = document.querySelector("input#object-of-problem").value;
})
