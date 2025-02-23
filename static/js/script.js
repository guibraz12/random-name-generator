document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("generate-name-button").addEventListener("click", fetchRandomName);
});

function fetchRandomName() {
    fetch('/random-name')
        .then(response => response.text())
        .then(name => {
            document.getElementById('random-name').innerText = name;
        })
        .catch(error => {
            console.error('Erro ao buscar o nome aleatório:', error);
        });
}