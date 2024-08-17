document.addEventListener('DOMContentLoaded', () => {
    const itemTableBody = document.getElementById('itemTableBody');
    itemTableBody.addEventListener('click', function(event) {
        if (event.target && event.target.nodeName === "BUTTON") {
            const action = event.target.getAttribute('data-action');
            const id = event.target.getAttribute('data-id');
            if (action === 'edit') {
                const name = event.target.getAttribute('data-name');
                const price = event.target.getAttribute('data-price');
                
                document.getElementById('id').value = id;
                document.getElementById('name').value = name;
                document.getElementById('price').value = price;

                const form = document.getElementById('itemForm');
                form.setAttribute('method', 'POST');
                const hiddenMethodInput = document.createElement('input');
                hiddenMethodInput.setAttribute('type', 'hidden');
                hiddenMethodInput.setAttribute('name', '_method');
                hiddenMethodInput.setAttribute('value', 'PUT');
                form.appendChild(hiddenMethodInput);
            }
        }
    });
});
