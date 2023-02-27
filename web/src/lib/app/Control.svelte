<script>
    import { PUBLIC_ECHO, PUBLIC_ECHO_PORT } from '$env/static/public'

	import Switch from "./Switch.svelte";
    import Button from "./Button.svelte";

    export let type;
    export let topic;
    export let options;
    export let title;

    // const publish = async (msg) => {
    //     const response = await fetch("http://localhost:1323/save");
    //     const data = await response.json();
    //     console.log(data);

    //     console.log(msg, topic)
    // }

    let result = null

    async function publish (msg) {
		const res = await fetch(`http://${PUBLIC_ECHO}:${PUBLIC_ECHO_PORT}/publishMessage`, {
			method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
            },
			body: JSON.stringify({
				"topic": topic,
                "message": msg
			})
		})
		
		const json = await res.json()
		result = JSON.stringify(json)
        console.log(result)
	}
    
</script>

<div class="p-5 bg-secondary rounded-lg flex ">
    {#if type == "switch"}
        <div>
            <h1 class="font-bold text-center pb-2">{title}</h1>
            <Switch publish={publish} options={options}/>
        </div>
    {/if}

    {#if type == "button"}
        <div>
            <h1 class="font-bold text-center pb-2">{title}</h1>
            <Button publish={publish} options={options}/>
        </div>
    {/if}
</div>
