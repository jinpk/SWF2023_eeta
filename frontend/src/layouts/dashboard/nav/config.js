// component
import SvgColor from '../../../components/svg-color';

// ----------------------------------------------------------------------

const icon = (name) => <SvgColor src={`/assets/icons/navbar/${name}.svg`} sx={{ width: 1, height: 1 }} />;

const navConfig = [
  {
    title: 'Ads Adaption',
    path: '/dashboard/app',
    icon: icon('ic_analytics'),
  },

  {
    title: '광고',
    path: '/dashboard/billboard',
    icon: icon('ic_cart'),
  },
  {
    title: 'STO',
    path: '/dashboard/sto',
    icon: icon('ic_cart'),
  },


  {
    title: '나의 자산',
    path: '/dashboard/wallet',
    icon: icon('ic_user'),
  },
  {
    title: '나의 광고',
    path: '/manage/billboard',
    icon: icon('ic_cart'),
  },
  {
    title: '나의 STO',
    path: '/manage/sto',
    icon: icon('ic_cart'),
  },
];

export default navConfig;
