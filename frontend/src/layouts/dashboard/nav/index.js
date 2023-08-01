import PropTypes from 'prop-types';
import { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
// @mui
import { styled, alpha } from '@mui/material/styles';
import { Box, Link, Drawer, Typography, Avatar } from '@mui/material';
// mock
import detectEthereumProvider from '@metamask/detect-provider'
import account from '../../../_mock/account';
// hooks
import useResponsive from '../../../hooks/useResponsive';
// components
import Logo from '../../../components/logo';
import Scrollbar from '../../../components/scrollbar';
import NavSection from '../../../components/nav-section';
//
import navConfig from './config';

// ----------------------------------------------------------------------

const NAV_WIDTH = 280;

export const StyledAccount = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  padding: theme.spacing(2, 2.5),
  borderRadius: Number(theme.shape.borderRadius) * 1.5,
  backgroundColor: alpha(theme.palette.grey[500], 0.12),
  overflow: 'hidden', // Clip the overflowing content
}));
// ----------------------------------------------------------------------

Nav.propTypes = {
  openNav: PropTypes.bool,
  onCloseNav: PropTypes.func,
};


export default function Nav({ openNav, onCloseNav }) {
  const { pathname } = useLocation();
  const [hasProvider, setHasProvider] = useState(null)
  const initialState = { accounts: [] }              
  const [wallet, setWallet] = useState(initialState) 

  useEffect(() => {
    const getProvider = async () => {
      const provider = await detectEthereumProvider({ silent: true })
      setHasProvider(Boolean(provider))
    }

    getProvider()
  }, [])

  const updateWallet = async (accounts) => {  
    setWallet({ accounts })                   
  }                                           

  const handleConnect = async () => {         
    const accounts = await window.ethereum.request({ 
      method: "eth_requestAccounts",                
    })                                              
    updateWallet(accounts)                          
  }                     

  const isDesktop = useResponsive('up', 'lg');

  useEffect(() => {
    if (openNav) {
      onCloseNav();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [pathname]);

  const renderContent = (
    <Scrollbar
      sx={{
        height: 1,
        '& .simplebar-content': { height: 1, display: 'flex', flexDirection: 'column' },
      }}
    >
      <Box sx={{ px: 2.5, py: 3, display: 'inline-flex' }}>
        <Logo />
      </Box>
      { hasProvider && wallet.accounts.length === 0 &&                               
        <Box sx={{ mb: 5, mx: 2.5 }} onClick={handleConnect}>
          <Link underline="none">
            <StyledAccount>
              <Avatar src={account.photoURL} alt="photoURL" />

              <Box sx={{ ml: 2 }}>
                <Typography variant="subtitle2" sx={{ color: 'text.primary' }}>
                  Connect Wallet
                </Typography>
              </Box>
            </StyledAccount>
          </Link>
        </Box>
      }
      
      { wallet.accounts.length > 0 &&                
        <Box sx={{ mb: 5, mx: 2.5 }}>
        <Link underline="none">
          <StyledAccount>
            <Avatar src={account.photoURL} alt="photoURL" />

            <Box sx={{ ml: 2 }}>
              <Typography variant="subtitle2" sx={{ color: 'text.primary' }}>
              {account.displayName}
              </Typography>

              <Typography variant="body2" sx={{ color: 'text.secondary', overflow: 'hidden', whiteSpace: 'nowrap', textOverflow: 'ellipsis', width: '140px' }}>
                {wallet.accounts[0]}
              </Typography>
            </Box>
          </StyledAccount>
        </Link>
        </Box>
      }

      <NavSection data={navConfig} />

      <Box sx={{ flexGrow: 1 }} />

    </Scrollbar>
  );

  return (
    <Box
      component="nav"
      sx={{
        flexShrink: { lg: 0 },
        width: { lg: NAV_WIDTH },
      }}
    >
      {isDesktop ? (
        <Drawer
          open
          variant="permanent"
          PaperProps={{
            sx: {
              width: NAV_WIDTH,
              bgcolor: 'background.default',
              borderRightStyle: 'dashed',
            },
          }}
        >
          {renderContent}
        </Drawer>
      ) : (
        <Drawer
          open={openNav}
          onClose={onCloseNav}
          ModalProps={{
            keepMounted: true,
          }}
          PaperProps={{
            sx: { width: NAV_WIDTH },
          }}
        >
          {renderContent}
        </Drawer>
      )}
    </Box>
  );
}
