import * as React from 'react';
import Box from '@mui/material/Box';

function helloWorld() {
  console.log()
}

export default function PodBox() {
  return (
    <Box
      sx={{
        width: '2.5rem',
        height: '2.5rem',
        margin: '0.5rem',
        bgcolor: 'primary.dark',
        '&:hover': {
          backgroundColor: 'primary.main',
          opacity: [0.9, 0.8, 0.7],
        }
      }}
      onClick={helloWorld}
    />
  )
}
