import React from "react";
import { Link, usePage } from "@inertiajs/react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from '@/components/ui/button';
import { Menu, Plus } from "lucide-react";
import { ModeToggle } from "@/components/ModeToggle";


const navigation = [
  { name: 'Home', href: '/' },
  { name: 'About', href: '/about' },
  { name: 'Projects', href: '#' },
  { name: 'Calendar', href: '#' },
]


const dropdownItems = [
  { name: 'Your Profile', href: '#' },
  { name: 'Settings', href: '#' },
  { name: 'Sign Out', href: '#' },
]

function classNames(...classes) {
  return classes.filter(Boolean).join(' ')
}

const Navbar = () => {
  const { url, component } = usePage();
  return (
    <nav className="shadow">
      <div className="mx-auto max-w-8xl px-4 sm:px-6 lg:px-8">
        <div className="flex h-16 justify-between">
          <div className="flex">
            <div className="-ml-2 mr-2 flex items-center md:hidden">
              <Button variant="outline" size="icon">
                <Menu />
              </Button>
            </div>
            <div className="flex shrink-0 items-center">
              <img className="h-8 w-auto" src="https://tailwindui.com/plus/img/logos/mark.svg?color=white" alt="Your Company" />
            </div>
            <div className="hidden md:ml-6 md:flex md:space-x-8">
              {navigation.map((item) => (
                <Link
                  key={item.name}
                  href={item.href}
                  aria-current={url === item.href ? 'page' : undefined}
                  className={classNames(
                    url === item.href ? 'border-white' : 'text-primary/90 border-transparent',
                    'inline-flex items-center border-b-2 px-1 pt-1 text-sm font-medium hover:border-primary/90',
                  )}
                >
                  {item.name}
                </Link>
              ))}
            </div>
          </div>
          <div className="flex items-center">
            <div className="shrink-0">
              <Button>
                <Plus />
                Add Todo
              </Button>
            </div>
            <div className="hidden md:ml-4 md:flex md:shrink-0 md:items-center">
              { /* Profile dropdown */}
              <div className="relative ml-3">
                <div>
                  <DropdownMenu>
                    <DropdownMenuTrigger>
                      <Avatar>
                        <AvatarImage src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80" />
                        <AvatarFallback>AS</AvatarFallback>
                      </Avatar>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent>
                      {dropdownItems.map((item) => (
                        <DropdownMenuItem key={item.name}><Link href={item.href}>{item.name}</Link></DropdownMenuItem>
                      ))}
                    </DropdownMenuContent>
                  </DropdownMenu>
                </div>
              </div>
              <div className="ml-3">
                <ModeToggle />
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="md:hidden">
        <div className="space-y-1 pb-3 pt-2">
          {navigation.map((item) => (
            <Link
              key={item.name}
              href={item.href}
              aria-current={url === item.href ? 'page' : undefined}
              className={classNames(
                url === item.href ? 'bg-secondary/50 border-primary text-primary' : 'text-primary/90 border-transparent',
                'block border-l-4 py-2 pl-3 pr-4 font-medium hover:border-primary/90 hover:bg-secondary/50 hover:text-primary/90 sm:pl-5 sm:pr-6',
              )}
            >
              {item.name}
            </Link>
          ))}
        </div>
        <div className="border-t border-gray-200 pb-3 pt-4">
          <div className="flex items-center px-4 sm:px-6">
            <div className="shrink-0">
              <Avatar>
                <AvatarImage src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80" />
                <AvatarFallback>AS</AvatarFallback>
              </Avatar>
            </div>
            <div className="ml-3">
              <div className="text-base font-medium text-primary">Tom Cook</div>
              <div className="text-sm font-medium text-primary">tom@example.com</div>
            </div>
          </div>
          <div className="mt-3 space-y-1">
            {dropdownItems.map((item) => (
              <Link
                key={item.name}
                href={item.href}
                className={
                  'block px-4 py-2 text-base font-medium text-primary hover:bg-secondary/50 hover:text-primary/90 sm:px-6'
                }
              >
                {item.name}
              </Link>
            ))}
          </div>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
