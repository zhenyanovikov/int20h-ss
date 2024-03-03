import { useState } from "react";
import { IconButton, Avatar, Menu, MenuItem, Box } from "@mui/material";
import { useTranslation } from "react-i18next";
import { getFullName } from "../../../helpers/user";
import { useLogout } from "../../../api/auth";
import { useGetUser } from "../../../api/user";

function UserMenu() {
  const { t } = useTranslation();

  const { data: userData, isLoading: isGetUserLoading } = useGetUser();
  const logout = useLogout();

  const [anchorEl, setAnchorEl] = useState(null);

  const isOpen = !!anchorEl;

  return (
    <Box>
      <IconButton onClick={handleOpenMenu}>
        {isGetUserLoading || !userData ? (
          <Avatar />
        ) : (
          <Avatar
            alt={getFullName(userData.firstName, userData.lastName)}
            src={userData.avatarUrl}
          />
        )}
      </IconButton>
      <Menu
        anchorEl={anchorEl}
        open={isOpen}
        onClick={handleCloseMenu}
        onClose={handleCloseMenu}
      >
        <MenuItem onClick={handleLogoutClick}>
          {t("molecules.userMenu.menu.logout")}
        </MenuItem>
      </Menu>
    </Box>
  );

  function handleOpenMenu(event) {
    setAnchorEl(event.currentTarget);
  }

  function handleCloseMenu() {
    setAnchorEl(null);
  }

  function handleLogoutClick() {
    logout();
  }
}

export default UserMenu;
