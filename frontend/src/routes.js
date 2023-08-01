import { Navigate, useRoutes } from 'react-router-dom';
// layouts
import DashboardLayout from './layouts/dashboard';
import SimpleLayout from './layouts/simple';
//
import LoginPage from './pages/LoginPage';
import Page404 from './pages/Page404';
import AddBillboardPage from './pages/AddBillboardPage';
import AddStoPage from './pages/AddStoPage';
import DashboardAppPage from './pages/DashboardPage';
import DepositPage from './pages/DepositPage';
import WalletPage from './pages/WalletPage';
import FundPage from './pages/FundPage';
import BillboardPage from './pages/BillboardPage';
import BidPage from './pages/BidPage';
import StoPage from './pages/StoPage';
import DetailBillboardPage from './pages/DetailBillboardPage';
import DetailStoPage from './pages/DetailStoPage';
import MyBillboardPage from './pages/MyBillboardPage';
import MyStoPage from './pages/MyStoPage';
import StoMarketPage from './pages/StoMarketPage';

// ----------------------------------------------------------------------

export default function Router() {
  const routes = useRoutes([
    {
      path: '/dashboard',
      element: <DashboardLayout />,
      children: [
        { element: <Navigate to="/dashboard/app" />, index: true },
        { path: 'app', element: <DashboardAppPage /> },
        { path: 'wallet', element: <WalletPage /> },
        { path: 'fund', element: <FundPage /> },
        { path: 'bid', element: <BidPage /> },
        { path: 'billboard', element: <BillboardPage /> },
        { path: 'sto', element: <StoPage /> },
        { path: 'market', element: <StoMarketPage /> },
        { path: 'billboard/:id', element: <DetailBillboardPage /> },
        { path: 'sto/:id', element: <DetailStoPage /> },
      ],
    },
    {
      path: '/manage',
      element: <DashboardLayout />,
      children: [
        { path: 'billboard', element: <MyBillboardPage /> },
        { path: 'sto', element: <MyStoPage /> },
      ],
    },
    { path: 'login', element: <LoginPage /> },
    { path: 'new-billboard', element: <AddBillboardPage /> },
    { path: 'new-sto', element: <AddStoPage /> },
    { path: 'deposit', element: <DepositPage /> },
    {
      element: <SimpleLayout />,
      children: [
        { element: <Navigate to="/dashboard/app" />, index: true },
        { path: '404', element: <Page404 /> },
        { path: '*', element: <Navigate to="/404" /> },
      ],
    },
    {
      path: '*',
      element: <Navigate to="/404" replace />,
    },
  ]);

  return routes;
}
