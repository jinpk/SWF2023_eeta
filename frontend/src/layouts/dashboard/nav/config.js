// component
import SvgColor from '../../../components/svg-color';

// ----------------------------------------------------------------------

const icon = (name) => <SvgColor src={`/assets/icons/navbar/${name}.svg`} sx={{ width: 1, height: 1 }} />;

const navConfig = [
  {
    title: '애즈 어댑션',
    path: '/dashboard/app',
    icon: icon('ic_analytics'),
  },

  {
    title: '광고 입찰하기',
    path: '/dashboard/billboard',
    icon: icon('ic_cart'),
  },
  {
    title: '모금중인 STO',
    path: '/dashboard/sto',
    icon: icon('ic_cart'),
  },
  {
    title: 'STO 거래소',
    path: '/dashboard/market',
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
