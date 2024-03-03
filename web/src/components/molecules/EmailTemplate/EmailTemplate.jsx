import {
  Card,
  CardActions,
  CardContent,
  Button,
  Stack,
  Typography,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { ROUTE } from "../../../constants/router";

function EmailTemplate({ emailTemplate }) {
  const { t } = useTranslation();
  const navigate = useNavigate();

  return (
    <Card>
      <CardContent>
        <Typography
          variant="h6"
          sx={{
            WebkitBoxOrient: "vertical",
            WebkitLineClamp: {
              xs: "2",
              md: "1",
            },
            display: "-webkit-box",
            overflow: "hidden",
          }}
        >
          {emailTemplate.subject}
        </Typography>
      </CardContent>
      <CardActions>
        <Stack flex={1} direction="row" justifyContent="flex-end">
          <Button onClick={handleUseTemplate}>
            {t("molecules.emailTemplate.button.use")}
          </Button>
        </Stack>
      </CardActions>
    </Card>
  );

  function handleUseTemplate() {
    navigate(ROUTE.SEND_EMAIL, { state: { emailTemplate } });
  }
}

export default EmailTemplate;
