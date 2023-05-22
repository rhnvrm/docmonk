import { ActionIcon, AppShell, Footer, Group, Header, Text, useMantineColorScheme } from "@mantine/core";
import { NavbarSearch } from "./Navbar";
import { IconMoonStars, IconSun } from "@tabler/icons-react";


export const ApplicationContainer = ({ children }: any) => {
    const { colorScheme, toggleColorScheme } = useMantineColorScheme();

    return (
        <AppShell
            padding="md"
            styles={(theme) => ({
                main: { backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.colors.gray[0] },
            })}
            header={
                <Header height={60} p="md">
                    <Group position="apart" spacing="xl">
                        <Text size="lg" weight="bolder">
                            Docmonk
                        </Text>
                        <ActionIcon
                            onClick={() => toggleColorScheme()}
                            size="lg"
                            sx={(theme) => ({
                                backgroundColor:
                                    theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
                                color: theme.colorScheme === 'dark' ? theme.colors.yellow[4] : theme.colors.blue[6],
                            })}
                        >
                            {colorScheme === 'dark' ? <IconSun size="1.2rem" /> : <IconMoonStars size="1.2rem" />}
                        </ActionIcon>
                    </Group>
                </Header>
            }
            navbar={
                <NavbarSearch></NavbarSearch>
            }
        >
            {children}
        </AppShell>
    )
}