<script>
	import { goto } from '$app/navigation';
	import Button, { Label } from '@smui/button';
</script>

<div class="container">
	<Button
		on:click={() => {
			goto('/login');
		}}
		variant="raised"
	>
		<Label>시작하기</Label>
	</Button>
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 100%;
	}
</style>
