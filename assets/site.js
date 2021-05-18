const Controller = {
    search: (ev) => {
        ev.preventDefault();
        const data = Object.fromEntries(new FormData(form));
        if (data.query.trim() == "") {
            alert("Please enter text to search!");
            return
        }
        document.getElementById("loading").style.display = "block";
        const response = fetch(`/search?q=${data.query}`).then((response) => {
            response.json().then((results) => {
                if (results.code == 200) {
                    if (results.data.lenth == 0) {
                        alert(`No result for ${data.query}`);
                        return
                    }
                } else {
                    alert("ERROR! ", message);
                    return
                }

                document.getElementById("loading").style.display = "none";
                Controller.updateList(results.data);
            });
        });
    },

    updateList: (results) => {
        const rows = [];
        if (results.length == 0) {
            rows.push(
                `
                <div class="list-group-item list-group-item-action border-none">
                    <div class="not-found">
                        <div><img src="/assets/not-found.png" /></div>
                        Sorry, the items you search is not found!
                    </div>
                    
                </div>
                `
            );
        }

        for (let result of results) {
            rows.push(
                `
                <div class="list-group-item list-group-item-action border-none">
                    <a href="#" class="title-item d-flex w-100 justify-content-between">
                        <h5 class="mb-1">${result.title}</h5>
                    </a>
                    <p class="mb-1">${result.content}</p>
                    <br />
                </div>
                `
            );
        }
        resultList.innerHTML = rows.join(" ");
    },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);