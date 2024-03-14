
import React, { useEffect, useState } from 'react';
import OutlinedCard from './card';
import PropTypes from 'prop-types';
import Toolbar from '@mui/material/Toolbar';
import useScrollTrigger from '@mui/material/useScrollTrigger';
import Box from '@mui/material/Box';
import Fab from '@mui/material/Fab';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import Fade from '@mui/material/Fade';

function ScrollTop(props) {
    const { children, window } = props;
    const trigger = useScrollTrigger({
      target: window ? window() : undefined,
      disableHysteresis: true,
      threshold: 100,
    });
  
    const handleClick = (event) => {
      const anchor = (event.target.ownerDocument || document).querySelector(
        '#back-to-top-anchor'
      );
  
      if (anchor) {
        anchor.scrollIntoView({
          block: 'center',
        });
      }
    };

    return (
        <Fade in={trigger}>
          <Box
            onClick={handleClick}
            role="presentation"
            sx={{ position: 'fixed', bottom: 16, right: 16 }}
          >
            {children}
          </Box>
        </Fade>
      );
}
ScrollTop.propTypes = {
    children: PropTypes.element.isRequired,
    /**
     * Injected by the documentation to work in an iframe.
     * You won't need it on your project.
     */
    window: PropTypes.func,
  };

export default function Blogs() {
    const [Actions, setActions] = useState([]);
    // const url = "https://swapi.dev/api/people/1"

    useEffect(() => {
      async function fetchData() {
          try {
              const res = await fetch("http://localhost:8080/SubscribedBlogs", {
                method : "POST",
              });

              if (!res.ok) {
                  throw new Error('Network response was not ok');
              }

              const result = await res.json();
              setActions(Object.entries(result));
              // setActions(result);

          } catch (error) {
              console.error('Error fetching data:', error.message);
          }
      }

      fetchData();
  }, []);

    useEffect(() => {
        // Now you can perform actions after the state has been updated
        console.log(Actions);
    }, [Actions]);

    return (
        <>
        <div className="blogs-wrapper">
              <Box>
                  <h1>All Recent Actions</h1>
                  <Toolbar id="back-to-top-anchor" />
                  {Actions.map(Action => (
                      // <Card blog={blog}></Card>
                      
                      <OutlinedCard blog={Action}></OutlinedCard>
                      
                  ))}
                  <ScrollTop {...Actions}>
                      <Fab size="small" aria-label="scroll back to top">
                      <KeyboardArrowUpIcon />
                      </Fab>
                  </ScrollTop>
              </Box>
              </div>
                
        </>
    );
}


            
