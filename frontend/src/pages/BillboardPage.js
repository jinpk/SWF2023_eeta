import { Helmet } from 'react-helmet-async';
// @mui
import { Grid, Container, Typography } from '@mui/material';
// components
import { BillboardCard, } from '../sections/@dashboard/billboard';
// mock
import POSTS from '../_mock/blog';

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
          {POSTS.map((post, index) => (
            <BillboardCard key={post.id} post={post} index={index} />
          ))}
        </Grid>
      </Container>
    </>
  );
}
