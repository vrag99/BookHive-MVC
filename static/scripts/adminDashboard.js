function goTo(link) {
    a = document.createElement("a");
    a.href = link;
    a.click();
}

async function addBook(button) {
    var { value: addedQuantity } = await Swal.fire({
        title: "Number of books to add",
        input: "number",
        showCancelButton: true,
        inputValidator: (value) => {
            if (value <= 0) {
                return "No. of books must be positive."
            }
        }
    })
    addedQuantity = parseInt(addedQuantity)

    if (isNaN(addedQuantity)) {
        goTo('/adminDashboard');
    } else {
        await axios
            .get("/adminDashboard", {
                params: {
                    id: button.id,
                    addedQuantity: addedQuantity,
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

async function removeBook(button) {
    var { value: removeQuantity } = await Swal.fire({
        title: "Number of books to remove",
        input: "number",
        showCancelButton: true,
        inputValidator: (value) => {
            value = parseInt(value)
            if (value < 0) {
                return "No. of books must be positive."
            } else if (button.dataset.available < value) {
                return "Can't remove more books than they exist."
            }
        }
    })
    removeQuantity = parseInt(removeQuantity)

    if (isNaN(removeQuantity)) {
        goTo('/adminDashboard');
    } else {
        await axios
            .get("/adminDashboard", {
                params: {
                    id: button.id,
                    removeQuantity: removeQuantity,
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

async function deleteBook(button) {
    await axios
        .get(`/adminDashboard/deleteBook/${button.id}`)
        .then(async (res) =>{
            await Swal.fire({
                title: "Deleted Successfully!",
                icon: "success",
                showConfirmButton: false,
                timer: 1000,
            });
        }).catch(async (err) =>{
            await Swal.fire({
                title: "Couldn't delete",
                text: "The book is already issued, requested for issue or to be returned",
                icon: "error",
                showConfirmButton: false,
                timer: 2000,
            });
        })
        goTo('/adminDashboard')
}