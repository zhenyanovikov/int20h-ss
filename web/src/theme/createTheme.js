import { createTheme as createMuiTheme } from "@mui/material";
import createPalette from "./createPalette";

function createTheme({ mode }) {
  const palette = createPalette({ mode });

  return createMuiTheme({
    palette,
  });
}

export default createTheme;
