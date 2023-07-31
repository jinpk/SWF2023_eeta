import PropTypes from 'prop-types';
// @mui
import { Grid } from '@mui/material';
import StoCard from './StoCard';

// ----------------------------------------------------------------------

StoList.propTypes = {
  products: PropTypes.array.isRequired,
};

export default function StoList({ products, ...other }) {
  return (
    <Grid container spacing={3} {...other}>
      {products.map((product) => (
        <Grid key={product.id} item xs={12} sm={6} md={3}>
          <StoCard product={product} />
        </Grid>
      ))}
    </Grid>
  );
}
