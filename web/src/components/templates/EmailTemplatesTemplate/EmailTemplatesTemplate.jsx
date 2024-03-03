import {
  Box,
  Button,
  CircularProgress,
  Container,
  Stack,
  Typography,
} from "@mui/material";
import { Create as CreateIcon } from "@mui/icons-material";
import { useTranslation } from "react-i18next";
import Empty from "../../../components/molecules/Empty/Empty";
import EmailTemplates from "../../organisms/EmailTemplates/EmailTemplates";
import { useNavigate } from "react-router-dom";
import { ROUTE } from "../../../constants/router";

function EmailTemplatesTemplate({
  emailTemplatesData,
  isGetEmailTemplatesLoading,
}) {
  const { t } = useTranslation();
  const navigate = useNavigate();

  return (
    <Box py={8}>
      <Container>
        <Stack spacing={4}>
          <Stack alignSelf="flex-start">
            <Button
              variant="contained"
              startIcon={<CreateIcon />}
              onClick={handleSendEmail}
            >
              {t("templates.emailTemplates.button.sendEmail")}
            </Button>
          </Stack>
          <Stack
            sx={{
              flexDirection: {
                xs: "column",
                md: "row",
              },
              alignItems: {
                xs: "flex-start",
                md: "center",
              },
              gap: 2,
            }}
          >
            <Typography variant="h4" component="h1">
              {t("templates.emailTemplates.title")}
            </Typography>
            <Box>
              <Button variant="outlined" onClick={handleCreateEmailTempate}>
                {t("templates.emailTemplates.button.createEmailTemplate")}
              </Button>
            </Box>
          </Stack>

          {isGetEmailTemplatesLoading || !emailTemplatesData ? (
            <Stack alignItems="center">
              <CircularProgress />
            </Stack>
          ) : (
            <>
              {!emailTemplatesData.length ? (
                <Empty />
              ) : (
                <EmailTemplates emailTemplates={emailTemplatesData} />
              )}
            </>
          )}
        </Stack>
      </Container>
    </Box>
  );

  function handleSendEmail() {
    navigate(ROUTE.SEND_EMAIL);
  }

  function handleCreateEmailTempate() {
    navigate(ROUTE.CREATE_EMAIL_TEMPLATE);
  }
}

export default EmailTemplatesTemplate;
