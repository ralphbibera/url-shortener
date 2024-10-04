<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	let count = 0;
</script>

<div>
	<h1>{count}</h1>
	<Button on:click={() => count++}>Increment</Button>
	<Button on:click={() => count--}>Decrement</Button>
</div>
