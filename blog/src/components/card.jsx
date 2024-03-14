import * as React from 'react';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { Link } from 'react-router-dom';
import SubscribeButton from './subscribe_btn';
import UnubscribeButton from './unsubcribe_btn';
import { useState } from 'react';

export default function OutlinedCard(props) {

    let len= props.blog[1][0].Blogs.title.length
    let blog = props.blog[1][0].Blogs
    // let Comment = props
    // console.log(props.user.email);
    
    return (
        <Box sx={{ minWidth: 275, width: 0.9, m: 'auto', mb: 2 }}>
            <Card variant="outlined">
                <CardContent>
                    <Typography variant="h5" component="div">
                        BlogID: {blog.id}
                    </Typography>
                    <Typography variant="h6" component="div">
                        Title: {blog.title.slice(3, len-4)}
                    </Typography>
                    <Typography color="text.secondary">
                        Author Handle: {blog.authorHandle}
                    </Typography>
                </CardContent>
                <CardActions>
                <Link to={{ pathname: `/blogs/${props.blog[0]}`, state: { blog: props.blog[1][0] } }}>
                    <Button size="small">Read More</Button>
                </Link>
                {/* <Button  size="small">Subscribe</Button> */}
                {!props.sub ? (
                <SubscribeButton email={props.user.email} blogId={props.blog[0]} />
                ) : (
                <UnubscribeButton email={props.user.email} blogId={props.blog[0]} />
                )}

                                



                </CardActions>
            </Card>
        </Box>
    );
}
