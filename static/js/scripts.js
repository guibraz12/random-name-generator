document.getElementById("generateNameButton").addEventListener("click", function() {
    fetch("/random-name")
        .then(response => response.text())
        .then(data => {
            document.getElementById("randomName").innerText = data;
        })
        .catch(error => console.error('Error:', error));
});