import { Box, Button, Container, Stack, Typography } from "@mui/material";

function LetteringTemplate({
  image,
  title,
  subtitle,
  buttonLabel,
  onButtonClick,
}) {
  const hasText = !!(title || subtitle);

  return (
    <Box py={8}>
      <Container>
        <Stack spacing={4} alignItems="center">
          {image}
          {hasText && (
            <Stack spacing={2} alignItems="center">
              <Typography variant="h3" component="h1" align="center">
                {title}
              </Typography>
              <Box
                sx={{
                  width: {
                    xs: "100%",
                    md: "80%",
                  },
                }}
              >
                <Typography variant="subtitle1" align="center">
                  {subtitle}
                </Typography>
              </Box>
            </Stack>
          )}
          {buttonLabel && (
            <Button variant="contained" onClick={onButtonClick}>
              {buttonLabel}
            </Button>
          )}
        </Stack>
      </Container>
    </Box>
  );
}

export default LetteringTemplate;
