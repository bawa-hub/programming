import Link from 'next/link';
import { useRouter } from 'next/router';
import React from 'react';

function SidebarMenu({ children, href }) {
  const router = useRouter();

  const active = router.asPath === href;

  return (
    <li className={active ? 'active' : 'inactive'}>
      <Link href={href}>
        <a>
          {children}
        </a>
      </Link>
    </li>
  );
}

export default SidebarMenu;
