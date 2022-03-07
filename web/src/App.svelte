<script lang="ts">
	interface User {
		id: number
		name: string
	}
	const fetchUsers = async ():Promise<User[]> => {
		 const res = await fetch(API_URL)
		 return res.json()
	}
	const promise = fetchUsers()
</script>

{#await promise}
  Loading
{:then users}
  {#each users as user}
    {user.id} : {user.name}<br/>
  {/each}
{:catch e}
 {e}
{/await}
