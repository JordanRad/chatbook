import React, { useState } from 'react';
import { Button, TextField, Typography } from '@mui/material';


const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        console.log(email, password)

        // Service call for login

        // Error handling

        // Eventually forward to chats if successful
    };

    return (
        <form onSubmit={handleSubmit}>
            <Typography variant="h4" gutterBottom>
                Login with credentials
            </Typography>
            <TextField
                label="Email"
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                fullWidth
                margin="normal"
                required
            />
            <TextField
                label="Password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                fullWidth
                margin="normal"
                required
            />
            <Button size='large' type="submit" variant="contained" color="primary">
                Login
            </Button>
        </form>
    );
};

export default Login;
