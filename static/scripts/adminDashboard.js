function goTo(link) {
    a = document.createElement("a");
    a.href = link;
    a.click();
}

async function addBook(btn) {
    var { value: addedQty } = await Swal.fire({
        title: "Number of books to add",
        input: "number",
        showCancelButton: true,
        inputValidator: (value) => {
            if (value <= 0) {
                return "No. of books must be positive."
            }
        }
    })
    addedQty = parseInt(addedQty)

    if (isNaN(addedQty)) {
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
                goTo('/adminDashboard');
            });
    }
}

async function removeBook(btn) {
    var {value: rmQty} = await Swal.fire({
        title: "Number of books to remove",
        input: "number",
        showCancelButton: true,
        inputValidator: (value) => {
            value = parseInt(value)
            if (value < 0) {
                return "No. of books must be positive."
            } else if (btn.dataset.available < value) {
                return "Can't remove more books than they exist."
            }
        }
    })
    rmQty = parseInt(rmQty)

    if (isNaN(rmQty)) {
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