import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import "../App.css";
import Comment from './Comment';

export default function Blog() {
    const [currentAction, setCurrentAction] = useState({});
    const { blogid } = useParams();

    useEffect(() => {
        async function fetchData() {
            try {
                const res = await fetch("http://localhost:8080/GroupedAction");

                if (!res.ok) {
                    throw new Error('Network response was not ok');
                }

                const result = await res.json();
                const resultArray = Object.entries(result);

                resultArray.forEach(([actionId, action]) => {
                    if (actionId === blogid && action.length > 0) {
                        setCurrentAction(action[0]);
                        // console.log(action[0])
                    }
                });
            } catch (error) {
                console.error('Error fetching data:', error.message);
            }
        }

        fetchData();
    }, [blogid]);

    // Check if currentAction.Blogs and currentAction.Comments are defined
    const removeTags = (str) => {
        return str.replace(/<\/?[^>]+(>|$)/g, '');
      };
    const blogTitle = currentAction.Blogs?.title || "Blog Title Not Available";
    const authorHandle = currentAction.Blogs?.authorHandle || "Author Not Available";

    return (
        <div>
            <div className='blog-content-container'>
                <h1 className='blog-h1'>{removeTags(blogTitle)}</h1>
                <p className='blog-p'>
                    {/* Add content here using currentAction.Blogs */}
                </p>
                <p className='blog-author'><span style={{ fontWeight: 700 }}>Author: </span>{authorHandle}</p>
                <div className="blog-comment-container">
                    {currentAction.Comments && currentAction.Comments.map((comment, index) => (
                        <Comment key={index} comment={comment} />
                    ))}
                </div>

            </div>
        </div>
    );
}
