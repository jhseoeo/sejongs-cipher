<script lang="ts">
	import '../global.css';
	import Button, { Label } from '@smui/button';
	import TopAppBar, { Row, Section, Title } from '@smui/top-app-bar';
</script>

<div class="app">
	<TopAppBar variant="static">
		<Row>
			<Section>
				<Title>Sejong's Cipher</Title>
			</Section>
			<Section align="end" toolbar>
				<Button color="secondary">
					<Label>Default</Label>
				</Button>
			</Section>
		</Row>
	</TopAppBar>
	<slot />
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		height: 100vh;
	}
</style>
