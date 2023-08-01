import PropTypes from 'prop-types';
// @mui
import { Grid } from '@mui/material';
import StoCard from './StoCard';

// ----------------------------------------------------------------------

StoList.propTypes = {
  stos: PropTypes.array.isRequired,
};

export default function StoList({ stos, ...other }) {
  return (
    <Grid container spacing={3} {...other}>
      {stos.map((sto) => (
        <Grid key={sto.id} item xs={12} sm={6} md={3}>
          <StoCard sto={sto} />
        </Grid>
      ))}
    </Grid>
  );
}
