import { IconButton } from "@mui/material";
import {
  LightMode as LightModeIcon,
  DarkMode as DarkModeIcon,
} from "@mui/icons-material";
import useClientStore from "../../../store/clientStore";
import { THEME_MODE } from "../../../constants/theme";

function ThemeToggleButton() {
  const themeMode = useClientStore((state) => state.themeMode);
  const toggleThemeMode = useClientStore((state) => state.toggleThemeMode);

  return (
    <IconButton onClick={toggleThemeMode} color="inherit">
      {themeMode === THEME_MODE.LIGHT ? <LightModeIcon /> : <DarkModeIcon />}
    </IconButton>
  );
}

export default ThemeToggleButton;
