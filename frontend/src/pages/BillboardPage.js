import { Helmet } from 'react-helmet-async';
// @mui
import { Grid, Container, Typography } from '@mui/material';
// components
import { BillboardCard, } from '../sections/@dashboard/billboard';
// mock
import BILLBOARD from '../_mock/billboard';

export default function BillboardPage() {
  return (
    <>
      <Helmet>
        <title> Dashboard: Blog | Minimal UI </title>
      </Helmet>

      <Container>
          <Typography variant="h4" mb={5}>
            Billboard
          </Typography>

        <Grid container spacing={3}>
          {BILLBOARD.map((billboard, index) => (
            <BillboardCard key={billboard.id} billboard={billboard} index={index} />
          ))}
        </Grid>
      </Container>
    </>
  );
}
