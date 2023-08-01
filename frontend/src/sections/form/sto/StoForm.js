import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Stack, TextField } from '@mui/material';
import { LoadingButton } from '@mui/lab';

export default function StoForm() {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    url: '',
    image: '',
  });

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleClick = () => {
    const apiUrl = 'http://ec2-3-37-36-76.ap-northeast-2.compute.amazonaws.com:8000/npos/';

    // Prepare the data object
    const data = {
      npo_name: formData.name,
      description: formData.description,
      website_url: formData.url,
      logo_url: formData.image,
    };

    // Send the data using fetch API
    fetch(apiUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'accept': 'application/json',
      },
      body: JSON.stringify(data),
    })
      .then((response) => {
        if (response.ok) {
          // Data sent successfully, navigate to the dashboard
          navigate('/dashboard', { replace: true });
        } else {
          // Handle errors here if needed
          console.error('Failed to send data.');
        }
      })
      .catch((error) => {
        console.error('Error occurred:', error);
      });
  };

  return (
    <>
      <Stack spacing={3} sx={{ my: 2 }}>
        <TextField
          name="name"
          label="Name"
          value={formData.name}
          onChange={handleChange}
        />
        <TextField
          name="description"
          label="Description"
          value={formData.description}
          onChange={handleChange}
        />
        <TextField
          name="url"
          label="Url"
          value={formData.url}
          onChange={handleChange}
        />
        <TextField
          name="image"
          label="Image"
          value={formData.image}
          onChange={handleChange}
        />
      </Stack>

      <LoadingButton fullWidth size="large" type="submit" variant="contained" onClick={handleClick}>
        Register
      </LoadingButton>
    </>
  );
}
