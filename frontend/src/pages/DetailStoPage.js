import { useState } from 'react';
import { Helmet } from 'react-helmet-async';
import { faker } from '@faker-js/faker';
import { useParams} from 'react-router-dom';
// @mui
import { useTheme } from '@mui/material/styles';
import { LoadingButton } from '@mui/lab';
import { Stack, TextField, Grid, Container, Typography, Button } from '@mui/material';
// components
import Iconify from '../components/iconify';
// sections
import {
  AppTasks,
  AppNewsUpdate,
  AppOrderTimeline,
  AppCurrentVisits,
  AppWebsiteVisits,
  AppTrafficBySite,
  AppWidgetSummary,
  AppCurrentSubject,
  AppConversionRates,
} from '../sections/@dashboard/app';

// ----------------------------------------------------------------------

export default function DetailStoPage() {
  const { id } = useParams();
  const theme = useTheme(); 
  const [formData, setFormData] = useState({
    amount: '',
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
      <Helmet>
        <title> Dashboard | Minimal UI </title>
      </Helmet>

      <Container maxWidth="xl">
        <Stack direction="row" alignItems="center" justifyContent="space-between" mb={5}>
          <Typography variant="h4" gutterBottom>
            STO
          </Typography>
        </Stack>
        <Typography variant="h4" gutterBottom>
          DDP A1 구역 전면광고
        </Typography>
        <Grid container spacing={3}>
          <Grid item xs={12} md={6} lg={4}>
            <Stack spacing={3} sx={{ my: 2 }}>
            <TextField
                name="amount"
                label="Amount"
                value={formData.amount}
                onChange={handleChange}
              />
            </Stack>

            <LoadingButton fullWidth size="large" type="submit" variant="contained" onClick={handleClick}>
              STO 참여하기
            </LoadingButton>
            <Typography>
            Available : 3,000 ST
            </Typography>
          </Grid>
          <Grid item xs={12} md={6} lg={4}>
            <AppOrderTimeline
              title="펀딩 일정"
              list={[...Array(4)].map((_, index) => ({
                id: faker.datatype.uuid(),
                title: [
                  '1983, orders, $4220',
                  '12 Invoices have been paid',
                  'Order #37745 from September',
                  'New order placed #XF-2356',
                ][index],
                type: `order${index + 1}`,
                time: faker.date.past(),
              }))}
            />
          </Grid>
          <Grid item xs={12} md={6} lg={4}>
            <AppTrafficBySite
              title="위험지수"
              list={[
                {
                  name: 'FaceBook',
                  value: 323234,
                  icon: <Iconify icon={'eva:facebook-fill'} color="#1877F2" width={32} />,
                },
                {
                  name: 'Google',
                  value: 341212,
                  icon: <Iconify icon={'eva:google-fill'} color="#DF3E30" width={32} />,
                },
                {
                  name: 'Linkedin',
                  value: 411213,
                  icon: <Iconify icon={'eva:linkedin-fill'} color="#006097" width={32} />,
                },
                {
                  name: 'Twitter',
                  value: 443232,
                  icon: <Iconify icon={'eva:twitter-fill'} color="#1C9CEA" width={32} />,
                },
              ]}
            />
          </Grid>

        </Grid>
      </Container>
    </>
  );
}
