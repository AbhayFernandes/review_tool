<script lang="ts">
    import '../../app.css';

    const handleClick = () => {
        var copyText = document.getElementById("commandView");

        // Select the text field

        // Copy the text inside the text field
        if (copyText?.textContent) {
            navigator.clipboard.writeText(copyText.textContent);
            // Alert the copied text
            // TODO: Make this a toast or something nicer
            alert("Copied the text: " + copyText.textContent);
            return {};
        } else {
            return {};
        }
    };

    // TODO: Update this for GRPC. 
    const load = async () => {
        const response = await fetch('http://localhost:8080/');
        const data = await response.json();
        return { props: { data } };
    }
</script>
  

<div class = "flex justify-center items-center min-h-screen bg-base-100">
    <div class="card bg-neutral text-neutral-content w-96">
      <div class="card-body items-center text-center">
        <h2 class="card-title">Login</h2>
        <p>Copy and run the following command to Authenticate yourself!</p>
        <div class="mockup-code">
            <pre data-prefix="$"><code id="commandView">TODO: Insert crev command here!</code></pre>
        </div>
        <div class="card-actions justify-end">
          <button on:click={handleClick} class="btn btn-primary-content">Copy</button>
        </div>
      </div>
    </div>
</div>
