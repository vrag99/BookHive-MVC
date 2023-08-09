function goTo(link) {
    a = document.createElement("a");
    a.href = link;
    a.click();
}

async function addBook(btn) {
    var addedQty = prompt("Enter the no. of books to add");
    if (addedQty <= 0) {
        alert("No. of books must be positive");
    } else if (isNaN(addedQty)) {
        goTo('/adminDashboard');
    } else {
        await axios
            .get("/adminDashboard", {
                params: {
                    id: btn.id,
                    addedQty: addedQty,
                },
            })
            .then(async (res) => {
                await Swal.fire({
                    title: "Added Successfully!",
                    icon: "success",
                    showConfirmButton: false,
                    timer: 1000,
                });
                goTo('/admindashboard');
            });
    }
}

async function removeBook(btn) {
    var rmQty = parseInt(prompt("Enter the no. of books to remove"));
    if (rmQty < 0) {
        alert("No. of books must be positive");
    } else if (btn.dataset.available < rmQty) {
        alert("Can't remove more books than they exist.");
    } else if (isNaN(rmQty)) {
        goTo('/adminDashboard');
    } else {
        await axios
            .get("/adminDashboard", {
                params: {
                    id: btn.id,
                    rmQty: rmQty,
                },
            })
            .then(async (res) => {
                await Swal.fire({
                    title: "Removed Successfully!",
                    icon: "success",
                    showConfirmButton: false,
                    timer: 1000,
                });
            })
            .catch(async (err) => {
                await Swal.fire({
                    title: "Clear pending requests first",
                    icon: "error",
                    showConfirmButton: false,
                    timer: 1000,
                });
            });
        goTo('/adminDashboard');
    }
}

var viewMode = document.getElementById("viewMode");
viewMode.addEventListener("change", () => {
    var mode = viewMode.value;
    if (mode != "all") goTo(`/adminDashboard/${mode}`);
    else goTo("/adminDashboard");
});