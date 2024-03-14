import React from 'react';
import Button from '@mui/material/Button';
const UnubscribeButton = (props) => {
    const handleSubscribe = async (blogId) => {
        try {
            const response = await fetch("http://localhost:8080/unsubscribe", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                }, 
                body: JSON.stringify({ email: props.email, blogid: Number(blogId) })
            });

            if (!response.ok) {
                throw new Error('Failed to subscribe');
            }

            // Handle successful subscription
            console.log('Unsubscribed successfully');
        } catch (error) {
            console.error('Error subscribing:', error.message);
        }
    };

    return (
        <Button onClick={() => handleSubscribe(props.blogId)} size="small">
            Unsubscribe
        </Button>
    );
};

export default UnubscribeButton;
