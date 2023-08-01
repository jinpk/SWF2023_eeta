import PropTypes from 'prop-types';
// @mui
import { Grid } from '@mui/material';
import BillboardCard from './BillboardCard';

// ----------------------------------------------------------------------

BillboardList.propTypes = {
  billboards: PropTypes.array.isRequired,
};

export default function BillboardList({ billboards, ...other }) {
  return (
    <Grid container spacing={3} {...other}>
      {billboards.map((billboard) => (
        <Grid key={billboard.id} item xs={12} sm={6} md={3}>
          <BillboardCard billboard={billboard} />
        </Grid>
      ))}
    </Grid>
  );
}
