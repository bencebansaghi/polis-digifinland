// Copyright (C) 2012-present, The Authors. This program is free software: you can redistribute it and/or  modify it under the terms of the GNU Affero General Public License, version 3, as published by the Free Software Foundation. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License for more details. You should have received a copy of the GNU Affero General Public License along with this program.  If not, see <http://www.gnu.org/licenses/>.

import React from "react";
import { useTranslation } from 'react-i18next';

const Footer = () => {
 const { t } = useTranslation();
 return (
    <footer style={{ position: 'absolute', left: '0'}}>
      <div style={{ display: 'flex', alignItems: 'flex-start', flexWrap: 'wrap', backgroundColor: '#003d6d', color: '#fff', fontSize: '0.875rem', position: 'relative', bottom: '0', paddingTop: '20px',paddingBottom: '20px' }}>
        <div style={{ width: '30%', minWidth: '300px', display: 'flex', flexDirection: 'column', marginLeft: '20px'}}>
            <h3>{t('footer.whatis')}</h3>
            <p>{t('footer.desc')}</p>
        </div>
        <div style={{ width: '30%', minWidth: '300px', display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
            <ul style={{ display: 'flex',flexDirection: 'column', listStyle: 'none',}}>
                <li style={{ margin: '10px'}}><a target="_blank" rel="noreferrer" style={{ color: '#FFF', border: 0 }} href="https://digifinland.fi/tietosuoja/">{t('footer.links.privacy')} 🡥</a></li>
                <li style={{ margin: '10px'}}><a target="_blank" rel="noreferrer" style={{ color: '#FFF', border: 0 }} href="https://digifinland.fi/toimintamme/polis-kansalaiskeskustelualusta/">{t('footer.links.info')} 🡥</a></li>
                <li style={{ margin: '10px'}}><a target="_blank" rel="noreferrer" style={{ color: '#FFF', border: 0 }} href="https://compdemocracy.org/Welcome/">The Computational Democracy Project 🡥</a></li>
                <li style={{ margin: '10px'}}><a target="_blank" rel="noreferrer" style={{ color: '#FFF', border: 0 }} href="https://github.com/compdemocracy/polis">{t('footer.links.source')} 🡥</a></li>
            </ul>
        </div>
        <div style={{ width: '30%', minWidth: '300px', display: 'flex', flexDirection: 'column'}}>
            <ul style={{ display: 'flex',flexDirection: 'column', listStyle: 'none',}}>
                <li><h3 style={{ margin: '0'}}>{t('footer.provider')}</h3></li>
                <li><p style={{ margin: '0'}}>DigiFinland Oy</p></li>
                <li><p style={{ margin: '0'}}>Kuntatalo, Toinen linja 14</p></li>
                <li><p style={{ margin: '0'}}>00180 Helsinki</p></li>
            </ul>
        </div>
      </div>
    </footer>
 )
};

export default Footer;