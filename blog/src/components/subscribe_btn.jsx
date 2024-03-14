import React from 'react';
import Button from '@mui/material/Button';
const SubscribeButton = (props) => {
    const handleSubscribe = async (blogId) => {
        try {
            const response = await fetch("http://localhost:8080/subscribe", {
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
            console.log('Subscribed successfully');
        } catch (error) {
            console.error('Error subscribing:', error.message);
        }
    };

    return (
        <Button onClick={() => handleSubscribe(props.blogId)} size="small">
            Subscribe
        </Button>
    );
};

export default SubscribeButton;
