let uploadedImages = [];

const toolbarOptions = [
    [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
    ['bold', 'italic', 'underline', 'strike'],        // toggled buttons
    ['link', 'image'],
    [{ 'list': 'ordered' }, { 'list': 'bullet' }],
    ['clean']                                         // remove formatting button
];

const quill = new Quill('#editor', {
    theme: 'snow',
    modules: {
        toolbar: {
            container: toolbarOptions,
            handlers: {
                image: function () {
                    const input = document.createElement('input');
                    input.setAttribute('type', 'file');
                    input.setAttribute('accept', 'image/*');
                    input.click();

                    input.onchange = async () => {
                        const file = input.files[0];
                        const formData = new FormData();
                        formData.append('image', file);

                        const response = await fetch('/upload-image', {
                            method: 'POST',
                            body: formData,
                        });
                        const data = await response.json();


                        const range = this.quill.getSelection();
                        this.quill.insertEmbed(range.index, 'image', data.data.imagePath);

                        // Track the uploaded image URL
                        uploadedImages.push(data.data.imagePath);
                    };
                },
            },
        },
    },
});


async function handlePublish() {
    const content = quill.root.innerHTML;

    // Extract image URLs from content
    const parser = new DOMParser();
    const doc = parser.parseFromString(content, 'text/html');
    const usedImages = Array.from(doc.querySelectorAll('img')).map(img => img.src);

    const unusedImages = uploadedImages.filter(url => !usedImages.includes(url));

    // Send unused images to the server to delete
    // await fetch('/remove-images', {
    // method: 'POST',
    // headers: {
    //     'Content-Type': 'application/json',
    // },
    // body: JSON.stringify({ images: unusedImages }),
    // });

    const titleInput = document.getElementById("title")
    const formData = new FormData()
    formData.append("title", titleInput.value)
    if(usedImages[0]){
        formData.append("bannerImg",usedImages[0])
    }
    formData.append("content", content)

    await fetch('/admin/add', {
        method: 'POST',
        body: formData,
    });
}

const publishButton = document.getElementById("publish")

publishButton.addEventListener('click', handlePublish)

