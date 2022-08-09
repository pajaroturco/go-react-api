import React, {useEffect, useState} from "react";

function App(){

    const [name, setName] = useState("")
    const [users, setUsers] = useState([])

    async function loadUsers(){
        const response = await fetch(  import.meta.env.VITE_API + "/users")
        const data = await response.json()
        setUsers(data.data)
    }

    useEffect (() => {
        loadUsers()
    },[])
    
    const handleSubmit = async (e) => {
        e.preventDefault()
        const response = await fetch(  import.meta.env.VITE_API + "/users",{
            method: "POST",
            headers: {
              "Content-Type":"application/json"
            },
            body: JSON.stringify({name})
        })
        const data = await response.json()
        loadUsers()
    }
    
    
    return (
        <div>
            <form onSubmit={handleSubmit}>
                <input type="name" placeholder="Coloca tu nombre" onChange={(e) => setName(e.target.value)}/>

                <button>Guardar</button>
            </form>

            {users.map(user => (
                <li key={user._id}>{user.name}</li>
            ))}
        </div>
    )
}

export default App