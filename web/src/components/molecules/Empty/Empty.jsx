import {
  Stack,
  // Box,
  Typography,
} from "@mui/material";
import { useTranslation } from "react-i18next";
// import EmptyIcon from "../../../assets/icons/icons8-empty-256.png";

function Empty() {
  const { t } = useTranslation();

  return (
    <Stack alignItems="center" spacing={1}>
      {/* <Box
        component="img"
        sx={{
          width: 128,
          height: 128,
        }}
        alt=""
        src={EmptyIcon}
      /> */}
      <Typography color="text.secondary">
        {t("molecules.empty.text.noData")}
      </Typography>
    </Stack>
  );
}

export default Empty;
